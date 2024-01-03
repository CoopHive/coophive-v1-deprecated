import tarfile
import tempfile
import requests
from fastapi import FastAPI
import gradio as gr
import time
import os
from pathlib import Path

app = FastAPI()


def run(module, inputs, request: gr.Request):
    if not request.query_params.get("userApiToken", "").startswith("lp-"):
        raise Exception("Please log in / register")
    res = requests.post("http://api/api/v1/jobs/sync", headers={
        "Authorization": "Bearer "+request.query_params["userApiToken"]
    }, json={
        "module": module,
        "inputs": inputs,
    })
    j = None
    try:
        j = res.json()
    except Exception as e:
        raise Exception(f"Error parsing JSON response: {res.text}")
    if j is not None:
        return getTarball(j["result"]["deal_id"])


def sdxl(prompt, request: gr.Request) -> Path:
    results_dir = run("sdxl:v0.9-lilypad1", {"PromptEnv": f"PROMPT={prompt}"}, request)
    return Path(results_dir+"/outputs/image-42.png")


def cowsay(message, request: gr.Request) -> str:
    results_dir = run("cowsay:v0.0.1", {"Message": message}, request)
    return open(results_dir+"/stdout").read()


def mistral7b(message, history, request: gr.Request):
    inst_str = ""
    for i, (q, a) in enumerate(history):
        inst_str += f"[INST]{q}[/INST]{a}"
        if i < len(history) - 1:
            inst_str += "\n"
    inst_str += f"[INST]{message}[/INST]"
    prompt = f"<s>{inst_str}"
    results_dir = run("mistral-7b-instruct:v0.1-lilypad3", {"PromptEnv": f"PROMPT={prompt}"}, request)
    return open(results_dir+"/stdout").read().split("[START]")[1].split("[/INST]")[-1]

# TODO: show the API call made to LilySaaS API in the UI, so users can see
# easily how to recreate it

# WHAT TO DO NEXT: make the functions above call the authenticated (for now)
# lilysaas API and actually work end-to-end

# If user is not logged in, make the frontend display a "log in to run this
# model" button with direct links to google and github OAuth

APPS = {
    "cowsay":
        gr.Interface(
            fn=cowsay,
            inputs=gr.Textbox(lines=2, placeholder="What would you like the cow to say?"),
            outputs="text",
            allow_flagging="never",
            css="footer {visibility: hidden}"
        ),
    "sdxl":
        gr.Interface(
            fn=sdxl,
            inputs=gr.Textbox(lines=2, placeholder="Enter prompt for Stable Diffusion XL"),
            outputs="image",
            allow_flagging="never",
            css="footer {visibility: hidden}"
        ),
    "mistral7b":
        gr.ChatInterface(mistral7b,
            css="footer {visibility: hidden}"
        ),
}

# should never access this route directly
@app.get("/")
def read_main():
    return {"message": "here be dragons"}

# download results from the solver to a tmp directory, returning the path
def getTarball(id: str) -> str:
    solverUrl = os.getenv("SOLVER_URL") + f"/api/v1/deals/{id}/files"

    r = requests.get(solverUrl, allow_redirects=True)

    with tempfile.TemporaryDirectory(delete=False) as temp_dir:
        temp_file_path = os.path.join(temp_dir, "temp.tar")
        with open(temp_file_path, "wb") as temp_file:
            temp_file.write(r.content)
            print(f" ---> HERE IS A FILE: {temp_file.name} from {solverUrl}")
        with tarfile.open(temp_file_path, "r") as tar:
            tar.extractall(temp_dir)
        return temp_dir



# must match path nginx/noxy is proxying to (see docker-compose.yml)
CUSTOM_PATH = "/gradio"

for (app_name, gradio_app) in APPS.items():
    print("mounting app", app_name, "->", gradio_app)
    app.mount(CUSTOM_PATH+"/"+app_name, gr.routes.App.create_app(gradio_app))
```
C:\Users\Jasper>conda env list

# conda environments:
#
# * -> active
# + -> frozen
base                     C:\Users\Jasper\miniconda3


C:\Users\Jasper>conda create --name langchain1.2 python=3.13.12
Do you accept the Terms of Service (ToS) for https://repo.anaconda.com/pkgs/main? [(a)ccept/(r)eject/(v)iew]: a
Do you accept the Terms of Service (ToS) for https://repo.anaconda.com/pkgs/r? [(a)ccept/(r)eject/(v)iew]: a
Do you accept the Terms of Service (ToS) for https://repo.anaconda.com/pkgs/msys2? [(a)ccept/(r)eject/(v)iew]: a
3 channel Terms of Service accepted
Retrieving notices: done
Channels:
 - defaults
Platform: win-64
Collecting package metadata (repodata.json): done
Solving environment: done

## Package Plan ##

  environment location: C:\Users\Jasper\miniconda3\envs\langchain1.2

  added / updated specs:
    - python=3.13.12


The following packages will be downloaded:

    package                    |            build
    ---------------------------|-----------------
    packaging-26.0             |  py313haa95532_0         196 KB
    pip-26.1.2                 |     pyhc872135_0         1.1 MB
    python-3.13.12             |h39c999c_100_cp313        15.9 MB
    python_abi-3.13            |          4_cp313           5 KB
    setuptools-82.0.1          |  py313haa95532_0         1.6 MB
    wheel-0.47.0               |  py313haa95532_0          97 KB
    ------------------------------------------------------------
                                           Total:        19.0 MB

The following NEW packages will be INSTALLED:

  bzip2              pkgs/main/win-64::bzip2-1.0.8-h2bbff1b_6
  ca-certificates    pkgs/main/win-64::ca-certificates-2026.5.14-haa95532_0
  libexpat           pkgs/main/win-64::libexpat-2.8.2-hd7fb8db_1
  libffi             pkgs/main/win-64::libffi-3.4.8-h42d73b9_3
done
#
# To activate this environment, use
#
#     $ conda activate langchain1.2
#
# To deactivate an active environment, use
#
#     $ conda deactivate


Channel "defaults" has the following notices:
  [info] -- Tue Jun  9 00:00:00 2026
  PyTorch 2.12 with CUDA support is now available to install with your current channel (Anaconda Main). Learn more: htts

WARNING conda.conda_pypi.main:notify_externally_managed_future(156):
  Did you know? You can install many PyPI packages with conda
  using the conda-pypi beta. Get started:
    https://docs.conda.io/projects/conda/en/stable/new-features.html


C:\Users\Jasper>conda env list

# conda environments:
#
# * -> active
# + -> frozen
base                     C:\Users\Jasper\miniconda3
langchain1.2             C:\Users\Jasper\miniconda3\envs\langchain1.2


C:\Users\Jasper>conda init
no change     C:\Users\Jasper\miniconda3\Scripts\conda.exe
no change     C:\Users\Jasper\miniconda3\Scripts\conda-script.py
no change     C:\Users\Jasper\miniconda3\condabin\conda.bat
no change     C:\Users\Jasper\miniconda3\Library\bin\conda.bat
no change     C:\Users\Jasper\miniconda3\condabin\_conda_activate.bat
no change     C:\Users\Jasper\miniconda3\condabin\rename_tmp.bat
no change     C:\Users\Jasper\miniconda3\condabin\conda_auto_activate.bat
no change     C:\Users\Jasper\miniconda3\condabin\conda_hook.bat
no change     C:\Users\Jasper\miniconda3\Scripts\activate.bat
no change     C:\Users\Jasper\miniconda3\condabin\activate.bat
no change     C:\Users\Jasper\miniconda3\condabin\deactivate.bat
modified      C:\Users\Jasper\miniconda3\Scripts\activate
modified      C:\Users\Jasper\miniconda3\Scripts\deactivate
modified      C:\Users\Jasper\miniconda3\etc\profile.d\conda.sh
modified      C:\Users\Jasper\miniconda3\etc\fish\conf.d\conda.fish
no change     C:\Users\Jasper\miniconda3\shell\condabin\Conda.psm1
modified      C:\Users\Jasper\miniconda3\shell\condabin\conda-hook.ps1
no change     C:\Users\Jasper\miniconda3\Lib\site-packages\xontrib\conda.xsh
modified      C:\Users\Jasper\miniconda3\etc\profile.d\conda.csh
modified      C:\Users\Jasper\Documents\WindowsPowerShell\profile.ps1
modified      HKEY_CURRENT_USER\Software\Microsoft\Command Processor\AutoRun

==> For changes to take effect, close and re-open your current shell. <==


C:\Users\Jasper>
```
```
pip install ipykernel
```
## 配置安装
```
# =========================================================
# LangChain 核心框架
# 作用：LangChain 1.x 主体、核心抽象、社区组件、实验性组件、文本切分器
# =========================================================
langchain==1.2.12
langchain-core==1.2.18
langchain-community==0.4.1
langchain-classic==1.0.2
langchain-text-splitters==1.1.1
langchain-experimental==0.4.1


# =========================================================
# jupyter
# 作用：交互式编程记事本，默认安装最新版即可
# =========================================================
jupyter


# =========================================================
# ollama
# 作用：调用本地大模型
# =========================================================
langchain-ollama==1.0.1
ollama==0.6.2


# =========================================================
# LangGraph / Agent 编排
# 作用：构建 Agent、状态图、多步骤工作流、检查点、PostgreSQL 持久化
# =========================================================
langgraph==1.1.2
langgraph-prebuilt==1.0.8
langgraph-checkpoint==4.0.1
langgraph-checkpoint-postgres==3.0.5
langgraph-sdk==0.3.9


# =========================================================
# MCP / FastMCP 集成
# 作用：构建 MCP Server、连接 MCP 工具、资源、Prompt
# =========================================================
mcp==1.27.0
fastmcp==3.2.4
langchain-mcp-adapters==0.2.1


# =========================================================
# 大模型供应商与 LangChain 适配器
# 作用：接入 OpenAI 协议、DeepSeek、Anthropic、OpenRouter、通义千问、腾讯等模型
# =========================================================
openai==2.26.0
anthropic==0.84.0
langchain-openai==1.1.11
langchain-deepseek==1.0.1
langchain-anthropic==1.3.4
langchain-openrouter==0.1.0
openrouter==0.7.11
dashscope==1.25.6
tencentcloud-sdk-python==3.1.86
PyJWT==2.10.1


# =========================================================
# 搜索工具 / 外部工具集成
# 作用：接入 Tavily 等联网搜索工具
# =========================================================
langchain-tavily==0.2.17


# =========================================================
# Web 服务 / API / SSE
# 作用：MCP HTTP 服务、FastAPI 服务、SSE 流式通信、本地 API 服务
# =========================================================
fastapi==0.135.1
uvicorn==0.46.0
sse-starlette==3.3.4
httpx==0.28.1
httpx-sse==0.4.3
aiohttp==3.12.14
requests==2.32.5
requests-toolbelt==1.0.0
websockets==16.0
watchfiles==1.1.1


# =========================================================
# 配置管理 / 数据校验 / 序列化
# 作用：读取 .env、Pydantic 配置、JSON/YAML、结构化输出
# =========================================================
python-dotenv==1.2.1
pydantic==2.12.5
pydantic-settings==2.12.0
PyYAML==6.0.3
orjson==3.11.7
jsonschema==4.26.0
jsonref==1.1.0
dataclasses-json==0.6.7


# =========================================================
# 日志 / 命令行 / 调试辅助
# 作用：日志输出、CLI、富文本终端输出、重试机制
# =========================================================
loguru==0.7.3
rich==14.3.3
typer==0.24.1
click==8.3.1
tenacity==9.1.4
tqdm==4.67.3
python-dateutil==2.9.0.post0
pytz==2026.2


# =========================================================
# 测试工具
# 作用：单元测试、教程代码验证
# =========================================================
pytest==9.0.3


# =========================================================
# 记忆相关
# 作用：PostgreSQL 检查点、SQL 数据源
# =========================================================
psycopg[binary]==3.3.3
psycopg-pool==3.3.0


# =========================================================
# RAG / 向量数据库 / 数据库连接
# 作用：Milvus 向量库、PostgreSQL 检查点、SQL 数据源
# =========================================================
langchain-milvus==0.3.3
pymilvus==2.6.12
SQLAlchemy==2.0.48


# =========================================================
# Embedding / Tokenizer / 文本处理
# 作用：TokenTextSplitter、SemanticChunker、文本相似度、传统 NLP 处理
# 注意：这里不包含 sentence-transformers，也不包含 torch
# =========================================================
tiktoken==0.12.0
numpy==2.4.4
scipy==1.17.1
scikit-learn==1.8.0
nltk==3.9.4
regex==2026.2.28
langdetect==1.0.9


# =========================================================
# HuggingFace / Transformers 基础组件
# 作用：本地模型、Tokenizer、部分文档解析模型可能会用到
# 注意：不包含 torch；如果加载本地深度学习模型，请单独安装匹配 CUDA 的 PyTorch
# =========================================================
transformers==5.3.0
tokenizers==0.22.2
huggingface-hub==1.11.0
safetensors==0.7.0


# =========================================================
# PyTorch的安装--cpu版本
# 作用：Unstructured的依赖
# =========================================================
torch==2.11.0
torchvision==0.26.0

# =========================================================
# Unstructured / LangChain Loader 文档解析
# 作用：PDF、Word、PPT、Excel、HTML、Markdown、图片文档等 Loader 支持
# 注意：unstructured-inference 可能依赖本地推理环境，torch/torchvision 请单独安装
# =========================================================
unstructured==0.20.6
unstructured-client==0.44.0
unstructured-inference==1.6.11
unstructured.pytesseract==0.3.15

pdfminer.six==20260107
pdf2image==1.17.0
pypdf==6.10.2
pypdfium2==5.8.0
pikepdf==10.5.1
pi-heif==1.3.0
pillow==12.2.0
opencv-python==4.13.0.92
onnx==1.21.0
onnxruntime==1.25.1

python-docx==1.2.0
python-pptx==1.0.2
openpyxl==3.1.5
xlrd==2.0.2
xlsxwriter==3.2.9
pandas==3.0.2

beautifulsoup4==4.14.3
html5lib==1.1
lxml==6.1.0
Markdown==3.10.2
jq==1.11.0
filetype==1.2.0
python-magic==0.4.27
python-iso639==2026.4.20
msoffcrypto-tool==6.0.0
python-oxmsg==0.0.2
olefile==0.47
pypandoc-binary==1.17
```
```
(langchain1.2) PS C:\workspace\langchain> pip install -r .\requirements.txt
Requirement already satisfied: langchain==1.2.12 in C:\Users\Jasper\miniconda3\envs\langchain1.2\Lib\site-packages (from -r .\requirements.txt (line 5)) (1.2.12)
Collecting langchain-core==1.2.18 (from -r .\requirements.txt (line 6))
  Downloading langchain_core-1.2.18-py3-none-any.whl.metadata (4.4 kB)
Collecting langchain-community==0.4.1 (from -r .\requirements.txt (line 7))
  Downloading langchain_community-0.4.1-py3-none-any.whl.metadata (3.0 kB)
Collecting langchain-classic==1.0.2 (from -r .\requirements.txt (line 8))
  Downloading langchain_classic-1.0.2-py3-none-any.whl.metadata (4.8 kB)
```
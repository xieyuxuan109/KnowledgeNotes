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
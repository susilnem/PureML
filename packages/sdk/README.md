[![PureML](/assets/SDKCoverImg.png)](https://pureml.com)

<h align="center">

# Track, version, compare and review your data and models.

</h>

# ⛳ Quick Access

<p align="center">
  <a
    href="https://pureml.mintlify.app"
  ><b>Documentation</b></a>
  &nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;
  <a
    href="https://www.youtube.com/watch?v=HdzLFEWS4s8&t=1s"
  ><b>Watch Demo</b></a>
  &nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;
  <a
    href="https://docs.pureml.com/docs/get-started/quickstart_tabular"
  ><b>Quick example</b></a>
  &nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;
  <a
    href="#"
  ><b>Get Instant Help</b></a>
  &nbsp;&nbsp;&nbsp; &nbsp;&nbsp;&nbsp;
  <a
    href="https://app.pureml.com/auth/signup"
  ><b>Sign Up for free</b></a>

</p>

</br>
</br>

<div align="center">
  <a
    href="https://pypi.org/project/pureml/"
  >
    <img alt="PyPi" src="https://img.shields.io/pypi/v/pureml?color=green&logo=pureml" />
  </a>
  &nbsp;
  <a
    href="https://python-poetry.org/"
  >
    <img src="https://img.shields.io/badge/poetry-1.1.14-blue?style=flat&logo=poetry&logoColor=white" />
  </a>
  &nbsp;
  <a
    href="https://opensource.org/licenses/Apache-2.0"
  >
    <img alt="License" src="https://img.shields.io/pypi/l/pureml?color=red&logo=Apache&logoColor=red" />
  </a>
  &nbsp;
  <a
    href="https://discord.gg/xNUHt9yguJ"
  >
    <img alt="Discord" src="https://img.shields.io/badge/Discord-Join%20Discord-blueviolet?style=flat&logo=discord&logoColor=white" />
  </a>
  &nbsp;
  <a
    href="https://pepy.tech/project/pureml"
  >
    <img alt="Downloads" src="https://static.pepy.tech/badge/pureml">
  </a>
  &nbsp;
  <a
    href="https://pypi.org/project/pureml/"
  >
    <img alt="^3.8" src="https://img.shields.io/pypi/pyversions/pureml">
  </a>
  &nbsp;
  <a
    href="https://pypi.org/project/pureml/"
  >
    <img alt="Coverage" src="https://img.shields.io/codecov/c/github/PureML-Inc/pureml">
  </a>
  &nbsp;
  <a
    href="https://pypi.org/project/pureml/"
  >
    <img alt="Coverage" src="https://img.shields.io/github/stars/pureml-inc/pureml?style=social">
  </a>

</div>

</br>
</br>

# 💎 Intro

PureML is an open-source version control for machine learning.

1. [Quick start](#quick-start)
1. [How it works](#how-it-works)
1. [Demo](#demo)
1. [Main Features](#main-features)
1. [Core design principles](#core-design-principles)
1. [Core abstractions](#core-abstractions)
1. [Why to get involved](#why-to-get-involved)
1. [Tutorials](#tutorials)

<br />

# ⏱ Quick start

You can install and run PureML using `pip`.

Install PureML

```bash
pip install pureml
```

<br />

# 📋 How it works

Just add a few lines of code. You don't need to change the way you work.

PureML is a Python library that uploads metadata to S3.

### Generating Data Lineage

1. Load Data

```python
@load_data(name='loading data')
def loading_data():

    return pd.read_csv('churn.csv')
```

2. Transform Data

```python
@transformer(name='fill missing values')
def fill_missing_values(df):
    return df.fillna()


@transformer(name='encode ordinal')
def encode_ordinal(df):
    col_ord = ['state', 'phone number']
    df_ord = df[col_ord]
    feat = OrdinalEncoder().fit_transform(df_ord)
    df[col_ord] = feat

    return df

@transformer(name='encode binary')
def encode_binary(df):

    df['voice mail plan'] = df['voice mail plan'].map({'yes':1, 'no':0})
    df['international plan'] = df['international plan'].map({'yes':1, 'no':0})
    df['churn'] = df['churn'].map({True:1, False:0})

    return df
```

3. Register Dataset

```python
@dataset(name='telecom churn', parent='encode binary')
def build_dataset():
    df = loading_data()

    df = fill_missing_values(df)

    df = encode_ordinal(df)

    df = encode_binary(df)

    return df

df = build_dataset()
```

This is how generated data lineage will look like in the UI

<h1 align="center">
    <img
      align="center"
      src="https://github.com/PureML-Inc/PureML/blob/main/assets/pipeline.png?raw=true"
      style="width:60%;"
    />
  </a>
</h1>

For more detailed explanation, please visit our [Documentation](https://docs.pureml.com)

# 💻 Demo

### Live demo

Build and run a PureML project to create data lineage and a model with our <b>[demo colab link](https://colab.research.google.com/drive/1LlrpaKiREwgesaRcnwkJP-w2MPesXf1t?usp=sharing)</b>.

### Demo video (2 min)

PureML quick start demo

[![PureML Demo Video](https://img.youtube.com/vi/HdzLFEWS4s8/0.jpg)](https://www.youtube.com/watch?v=HdzLFEWS4s8 "PureML Demo Video")

<sub><i>Click the image to play video</i></sub>

<br />

# 📍 [Main Features](https://docs.pureml.com/)

|                          |                                                                               |
| ------------------------ | ----------------------------------------------------------------------------- |
| Data Lineage             | Automatic generation of data lineage                                          |
| Dataset Versioning       | Object-based Automatic Semantic Versioning of datasets                        |
| Model Versioning         | Object-based Automatic Semantic Versioning of models                          |
| Comparision              | Comparing different versions of models or datasets                            |
| Branches (_Coming Soon_) | Separation between experimentation and production ready models using branches |
| Review (_Coming Soon_)   | Review and approve models, and datasets to production ready branch            |

<br />

# 🔮 Core design principles

|                                     |                                                                                                                 |
| ----------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| Easy developer experience           | An intuitive open source package aimed to bridge the gaps in data science teams                                 |
| Engineering best practices built-in | Integrating PureML functionalities in your code does not disrupt your workflow                                  |
| Object Versioning                   | A reliable object versioning mechanism to track changes to your datasets, and models                            |
| Data is a first-class citizen       | Your data is secure. It will never leave your system.                                                           |
| Reduce Friction                     | Have access to operations performed on data using data lineage without having to spend time on lengthy meetings |

<br />

# ⚙ Core abstractions

These are the fundamental concepts that PureML uses to operate.

|                                                                     |                                                                                                                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [Project](https://docs.pureml.com/docs/projects/about_projects)     | A data science project. This is where you store datasets, models, and their related objects. It is similar to a github repository with object storage.  |
| [Lineage](https://docs.pureml.com/docs/data/register_data_pipeline) | Contains a series of transformations performed on data to generate a dataset.                                                                           |
| Data Versioning                                                     | Versioning of the data should be comprehensible to the user and should encapsulate the changes in the data, its creation mechanism, among others.       |
| Model Versioning                                                    | Versioning of the model should be comprehensible to the user and should encapuslate the changes in training data, model architecture, hyper parameters. |
| Fetch                                                               | This functionality is used to fetch registered Models, and Datasets.                                                                                    |

<br />

# 🤝 Why to get involved

Version control is much more common in software than in machine learning. So why isn’t everyone using Git? Git doesn’t work well with machine learning. It can’t handle large files, it can’t handle key/value metadata like metrics, and it can’t record information automatically from inside a training script.

GitHub wasn’t designed with data as a core project component. This along with a number of other differences between AI and more traditional software projects makes GitHub a bad fit for artificial intelligence, contributing to the reproducibility crisis in machine learning.

From manually tracking models to git based versioning systems that do not follow an intuitive versioning mechanism, there is no standardized way to track objects. Using these mechanisms, it is hard enough to track or get your model from a month ago running, let alone of a teammates!

We are trying to build a version control system for machine learning objects. A mechanism that is object dependant and intuitive for users.

Lets build this together. If you have faced this issue or have worked out a similar solution for yourself, please join us to help build a better system for everyone.

<br />

# 🧮 Tutorials

- [Registering Data lineage](https://docs.pureml.com/docs/data/register_data_pipeline)
- [Registering models](https://docs.pureml.com/docs/models/register_models)
- [Quick Start: Tabular](https://docs.pureml.com/docs/get-started/quickstart_tabular)
- [Quick Start: Computer Vision](https://docs.pureml.com/docs/get-started/quickstart_cv)
- [Quick Start: NLP](https://docs.pureml.com/docs/get-started/quickstart_nlp)
- [Logging](https://docs.pureml.com/docs/log/overview)

<br />

# 🐞 Reporting Bugs

To report any bugs you have faced while using PureML package, please

1. Report it in [Discord](https://discord.gg/xNUHt9yguJ) channel
1. Open an [issue](https://github.com/PureML-Inc/PureML/issues)

<br />

# ⌨ Contributing and Developing

Lets work together to improve the features for everyone. For more details, please look at out [Contributing Guide](./CONTRIBUTING.md)

Work with mutual respect.

<br />

# 👨‍👩‍👧‍👦 Community

To get quick updates, feature release for PureML follow us on
|                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [<img alt="Twitter" height="20" src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white" />](https://twitter.com/getPureML)                   |
| [<img alt="LinkedIn" height="20" src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white" />](https://www.linkedin.com/company/pureml-inc/) |
| [<img alt="GitHub" height="20" src="https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white" />](https://github.com/PureML-Inc/PureML)               |
| [<img alt="GitHub" height="20" src="https://img.shields.io/badge/Discord-5865F2?style=for-the-badge&logo=discord&logoColor=white" />](https://discord.gg/DBvedzGu)                      |

# 📄 License

See the [Apache-2.0](./License) file for licensing information.

<br />

**WARNING: this model is intentionally vulnerable to a deserialization attack**
The model is trained on a diabetes dataset, and predicts whether a person has diabetes or not.
The dataset can be found here: [Link to PIMA Indian diabetes dataset.](https://www.kaggle.com/datasets/uciml/pima-indians-diabetes-database)
The model then has the following code injected in it so that we can scan it:

```python
command = "system"
malicious_code = """cat ~/.aws/secrets
    """
```
The example is taken from modelscan's examples [here](https://github.com/protectai/modelscan/blob/main/notebooks/xgboost_diabetes_classification.ipynb)

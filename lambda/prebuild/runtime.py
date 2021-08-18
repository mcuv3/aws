import importlib
import os

try:
    module = importlib.import_module(os.environ.get("HANDLER"))
    data = os.environ.get("EVENT_DATA",{})
    module.hello(data)
except Exception as e:
    print("Could not find path,",e)

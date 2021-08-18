import importlib
import os
from inspect import signature

try:
    handler = os.environ.get("HANDLER")
    mod = handler.split(".")
    func = mod.pop()
    path = ".".join(mod)
    module = importlib.import_module(path) # code.hello
    data = os.environ.get("EVENT_DATA",{})
    func = getattr(module, func)
    sig = signature(func)
    if len(sig.parameters) == 1:
        func(data)
    else:
        func()
except Exception as e:
    print("Could not find path,",e)

from django.shortcuts import HttpResponse
from django.http import JsonResponse
from django.views import View
import json
import ollama
from asgiref.sync import async_to_sync
from threading import Thread
from django.views.decorators.csrf import csrf_exempt
from django.utils.decorators import method_decorator
import random
# Create your views here.
@method_decorator(csrf_exempt, name='dispatch')
class chatView(View):
    def post(self, request, *args, **kwargs):
        modelExists = checkmodel()
        if modelExists == False:
            return HttpResponse("Server Failure! The model is not available at the moment!")
        data = json.loads(request.body)
        message = data.get("message")
        def generate(prompt1):
            nonlocal result
            result1 = ollama.chat(model="llama2-uncensored", messages=[{"role": "user", "content": prompt1}])
            result=result1["message"]["content"]
            #result = ollama.generate(model="llama2-uncensored", prompt= prompt1)
            #result["response"]
        result = {}
        thread = Thread(target=generate(message))
        thread.start()
        thread.join()
        return JsonResponse({"response": result})    

def checkmodel():
    modelExists=False
    if "llama2-uncensored" not in str(ollama.list()):
        modelExists=False
    else:
        modelExists=True
    return modelExists

def llama(request):
    return HttpResponse("Welcome to Llama chat. Send post request to path /chat. You can ask it about the objects at our store and various categories - and feel free to do so")

def chatstart(request):
    outputs=["Welcome to VidGame Shop. How could I help You today?", "Hello, glad to see you at VidGame Shop. How could I be of service?", "Welcome, I'm VidGame Shop Bot. How could I assist You?", "Hi, I'm VidGame Shop bot. Is there anything I can help You with?", "This is VidGame Shop Bot speaking. How can I be of service?"]
    pick=random.randint(0, len(outputs)-1)
    result=outputs[pick]
    return JsonResponse({"response": result})

def chatend(request):
    outputs=["Thank You for chatting. I hope I was able to help", "Thanks for the conversation and I hope to see You again soon", "It was great talking to You, take care.", "Glad I got to chat with You! I hope we'll get to talk some other time.", "Thank You for Your interest in VidGame Shop, feel free to come back anytime if You have any more questions!"]
    pick=random.randint(0, len(outputs)-1)
    result=outputs[pick]
    return JsonResponse({"response": result})
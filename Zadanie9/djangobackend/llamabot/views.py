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
import logging
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
            prompt0="Say whether the following phrase concerns video game shop's information and details, game platforms, available games or matters unrelated to video game shop using just one of the following words as the answer: details, platforms, games, other. Phrase: " + prompt1
            result0=ollama.generate(model="llama2-uncensored", prompt=prompt0)
            logging.warning(result0["response"])
            #options -> details, available platforms, available stock, other
            if "games" in str(result0["response"]).lower():
                game1=["Death Stranding", "PS5", 29.99, "Story-rich shooter"]
                game2=["Tom Clancy's Ghost Recon", "PC", 5.99, "Tactical shooter"]
                game3=["The Shore", "PC", 7.99, "Horror"]
                game4=["Tom Clancy's Splinter Cell", "PC", 5.99, "Stealth shooter"]
                game5=["Rayman Legends", "PS3", 7.99, "2D Platformer"]
                game6=["Ratchet & Clank", "PS4", 12.99, "3D Platformer shooter"]
                game7=["Avicii Invector", "PS4", 11.99, "Rhythm Game"]
                games=[game1, game2, game3, game4, game5, game6, game7]
                prompt="Answer the following phrase speaking directly to the user, taking into account relevant aspects from this information, showing game name, platform, price and genre:" + str(games) + ". In case of requiring specific game, platform, genre or price, display only games satisfying user's requirements. Phrase: " + prompt1
                result1=ollama.generate(model="llama2-uncensored", prompt=prompt)
            elif "platforms" in str(result0["response"]).lower():
                prompt="Answer the following phrase speaking directly to the user, taking into account that the VidGame Shop sells hardware and software for PC, PS3, PS4 and PS5 platforms. The phrase: " + prompt1
                result1=ollama.generate(model="llama2-uncensored", prompt=prompt)
            elif "details" in str(result0["response"]).lower():
                prompt="Answer the following phrase speaking directly to the user, taking into account important aspects of the following context: name - The VidGame Shop, online shop at http://localhost:8000/, established 2001, selling both new and old games. Phrase to answer: " + prompt1
                result1=ollama.generate(model="llama2-uncensored", prompt=prompt)
            else:
                result1=ollama.generate(model="llama2-uncensored", prompt="Say that you are sorry, but you do not feel like you'll be able to answer this specific question at the moment. Say you may have misundrestood the question if it was related to the shop, and ask the user to rephrase if it was the case. Be polite")
            #result1 = ollama.chat(model="llama2-uncensored", messages=[{"role": "user", "content": prompt1}])
            result=result1["response"]
            #result = ollama.generate(model="llama2-uncensored", prompt= prompt1)
            #result["response"]
        result = {}
        thread = Thread(target=generate(message))
        thread.start()
        thread.join()
        return JsonResponse({"response": result})    
#7B model, regardless of prompt manipulation may not work as expected at times...

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
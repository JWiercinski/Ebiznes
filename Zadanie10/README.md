GitHub Actions zebrane odpowiednio z repozytoriów gocloud oraz reactcloud

gocloud:

**master_jwgoserver.yml** - wstępna próba budowy aplikacji go, wysyłka maila potwierdzającego sukces, budowa i push dockerowego obrazu oraz deployment na chmurę

**master_jwclientreact.yml** - wstępna próba budowy aplikacji react, wysyłka maila potwierdzającego sukces, budowa i push dockerowego obrazu oraz deployment na chmurę

Oba workflowy bazują na zasadzie kontynuacji wyłącznie w przypadku sukcesu - porażka w jakiejkolwiek fazie anuluje dalsze działanie

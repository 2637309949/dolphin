function onOpenSwaggerUI(){
    window.open(chrome.extension.getURL('index.html'));
}

document.getElementById('go_backend_page').onclick=onOpenSwaggerUI
package app

import (
    "send-email/internal/config"
    "send-email/internal/handler"
    "send-email/internal/repository"
    "send-email/internal/service"
    "net/http"
)

func Run() {
    // Carregar configurações
    cfg := config.Load()

    // Inicializar dependências
    repo := repository.NewRepository(cfg)
    srv := service.NewService(repo)
    h := handler.NewHandler(srv)

    // Configurar rotas e iniciar o servidor
    http.HandleFunc("/campaign", h.CreateCampaign)
    http.ListenAndServe(cfg.ServerAddress, nil)
}

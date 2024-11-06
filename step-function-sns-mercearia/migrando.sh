curl -X 'POST' \
  'http://localhost:5226/api/Pedido/processar-pedidos' \
  -H 'accept: */*' \
  -H 'Content-Type: application/json' \
  -d '{
    "pedidos": [
      {
        "pedidoId": 1,
        "produtos": [
          {"produtoId": "PS5", "dimensoes": {"altura": 40, "largura": 10, "comprimento": 25}},
          {"produtoId": "Volante", "dimensoes": {"altura": 40, "largura": 30, "comprimento": 30}}
        ]
      },
      {
        "pedidoId": 2,
        "produtos": [
          {"produtoId": "Joystick", "dimensoes": {"altura": 15, "largura": 20, "comprimento": 10}},
          {"produtoId": "Fifa 24", "dimensoes": {"altura": 10, "largura": 30, "comprimento": 10}},
          {"produtoId": "Call of Duty", "dimensoes": {"altura": 30, "largura": 15, "comprimento": 10}}
        ]
      },
      {
        "pedidoId": 3,
        "produtos": [
          {"produtoId": "Headset", "dimensoes": {"altura": 25, "largura": 15, "comprimento": 20}}
        ]
      },
      {
        "pedidoId": 4,
        "produtos": [
          {"produtoId": "Mouse Gamer", "dimensoes": {"altura": 5, "largura": 8, "comprimento": 12}},
          {"produtoId": "Teclado Mecânico", "dimensoes": {"altura": 4, "largura": 45, "comprimento": 15}}
        ]
      },
      {
        "pedidoId": 5,
        "produtos": [
          {"produtoId": "Cadeira Gamer", "dimensoes": {"altura": 120, "largura": 60, "comprimento": 70}}
        ]
      },
      {
        "pedidoId": 6,
        "produtos": [
          {"produtoId": "Webcam", "dimensoes": {"altura": 7, "largura": 10, "comprimento": 5}},
          {"produtoId": "Microfone", "dimensoes": {"altura": 25, "largura": 10, "comprimento": 10}},
          {"produtoId": "Monitor", "dimensoes": {"altura": 50, "largura": 60, "comprimento": 20}},
          {"produtoId": "Notebook", "dimensoes": {"altura": 2, "largura": 35, "comprimento": 25}}
        ]
      },
      {
        "pedidoId": 7,
        "produtos": [
          {"produtoId": "Jogo de Cabos", "dimensoes": {"altura": 5, "largura": 15, "comprimento": 10}}
        ]
      },
      {
        "pedidoId": 8,
        "produtos": [
          {"produtoId": "Controle Xbox", "dimensoes": {"altura": 10, "largura": 15, "comprimento": 10}},
          {"produtoId": "Carregador", "dimensoes": {"altura": 3, "largura": 8, "comprimento": 8}}
        ]
      },
      {
        "pedidoId": 9,
        "produtos": [
          {"produtoId": "Tablet", "dimensoes": {"altura": 1, "largura": 25, "comprimento": 17}}
        ]
      },
      {
        "pedidoId": 10,
        "produtos": [
          {"produtoId": "HD Externo", "dimensoes": {"altura": 2, "largura": 8, "comprimento": 12}},
          {"produtoId": "Pendrive", "dimensoes": {"altura": 1, "largura": 2, "comprimento": 5}}
        ]
      }
    ]
  }'



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


exit 

curl -X 'POST' \
  'http://localhost:5166/api/Pedido/processar-pedidos' \
  -H 'accept: */*' \
  -H 'Content-Type: application/json' \
  -d '{
    "pedidos": [
      {
        "pedidoId": 1,
        "produtos": [
          {
            "produtoId": "PS5",
            "dimensoes": {
              "altura": 40,
              "largura": 10,
              "comprimento": 25
            }
          },
          {
            "produtoId": "Volante",
            "dimensoes": {
              "altura": 40,
              "largura": 30,
              "comprimento": 30
            }
          }
        ]
      },
      {
        "pedidoId": 2,
        "produtos": [
          {
            "produtoId": "Joystick",
            "dimensoes": {
              "altura": 15,
              "largura": 20,
              "comprimento": 10
            }
          },
          {
            "produtoId": "Fifa 24",
            "dimensoes": {
              "altura": 10,
              "largura": 30,
              "comprimento": 10
            }
          }
        ]
      }
    ]
  }'


exit


curl -X 'POST' \
  'http://localhost:5166/api/Pedido/processar-pedidos' \
  -H 'accept: */*' \
  -H 'Content-Type: application/json' \
  -d '{
    "pedidos": [
      {
        "pedido_id": 1,
        "produtos": [
          {"produto_id": "PS5", "dimensoes": {"altura": 40, "largura": 10, "comprimento": 25}},
          {"produto_id": "Volante", "dimensoes": {"altura": 40, "largura": 30, "comprimento": 30}}
        ]
      },
      {
        "pedido_id": 2,
        "produtos": [
          {"produto_id": "Joystick", "dimensoes": {"altura": 15, "largura": 20, "comprimento": 10}},
          {"produto_id": "Fifa 24", "dimensoes": {"altura": 10, "largura": 30, "comprimento": 10}},
          {"produto_id": "Call of Duty", "dimensoes": {"altura": 30, "largura": 15, "comprimento": 10}}
        ]
      },
      {
        "pedido_id": 3,
        "produtos": [
          {"produto_id": "Headset", "dimensoes": {"altura": 25, "largura": 15, "comprimento": 20}}
        ]
      },
      {
        "pedido_id": 4,
        "produtos": [
          {"produto_id": "Mouse Gamer", "dimensoes": {"altura": 5, "largura": 8, "comprimento": 12}},
          {"produto_id": "Teclado Mecânico", "dimensoes": {"altura": 4, "largura": 45, "comprimento": 15}}
        ]
      },
      {
        "pedido_id": 5,
        "produtos": [
          {"produto_id": "Cadeira Gamer", "dimensoes": {"altura": 120, "largura": 60, "comprimento": 70}}
        ]
      },
      {
        "pedido_id": 6,
        "produtos": [
          {"produto_id": "Webcam", "dimensoes": {"altura": 7, "largura": 10, "comprimento": 5}},
          {"produto_id": "Microfone", "dimensoes": {"altura": 25, "largura": 10, "comprimento": 10}},
          {"produto_id": "Monitor", "dimensoes": {"altura": 50, "largura": 60, "comprimento": 20}},
          {"produto_id": "Notebook", "dimensoes": {"altura": 2, "largura": 35, "comprimento": 25}}
        ]
      },
      {
        "pedido_id": 7,
        "produtos": [
          {"produto_id": "Jogo de Cabos", "dimensoes": {"altura": 5, "largura": 15, "comprimento": 10}}
        ]
      },
      {
        "pedido_id": 8,
        "produtos": [
          {"produto_id": "Controle Xbox", "dimensoes": {"altura": 10, "largura": 15, "comprimento": 10}},
          {"produto_id": "Carregador", "dimensoes": {"altura": 3, "largura": 8, "comprimento": 8}}
        ]
      },
      {
        "pedido_id": 9,
        "produtos": [
          {"produto_id": "Tablet", "dimensoes": {"altura": 1, "largura": 25, "comprimento": 17}}
        ]
      },
      {
        "pedido_id": 10,
        "produtos": [
          {"produto_id": "HD Externo", "dimensoes": {"altura": 2, "largura": 8, "comprimento": 12}},
          {"produto_id": "Pendrive", "dimensoes": {"altura": 1, "largura": 2, "comprimento": 5}}
        ]
      }
    ]
  }'

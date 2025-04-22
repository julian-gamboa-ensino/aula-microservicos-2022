const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const path = require('path');
const { PrismaClient } = require('@prisma/client');

const prisma = new PrismaClient();
const app = express();

// Middleware
app.use(bodyParser.json());
app.use(cors());

const FrontEnd = process.env.FrontEnd || 'FrontEnd2024';

// Servir arquivos estáticos do Angular
app.use(express.static(path.join(__dirname, FrontEnd)));



// Rota para obter todos os produtos
app.get('/api/products', async (req, res) => {
  try {
    const products = await prisma.product.findMany();
    res.json(products);
  } catch (error) {
    //res.status(500).json({ error: error.message });
    next(error); // Passa o erro para o middleware de tratamento de erros

  }
});

// Rota para criar um novo produto
app.post('/api/products', async (req, res) => {
  try {
    const product = await prisma.product.create({
      data: req.body,
    });
    res.status(201).json(product);
  } catch (error) {
    //res.status(400).json({ error: error.message });
    next(error); // Passa o erro para o middleware de tratamento de erros

  }
});

// Rota para obter um produto pelo ID
app.get('/api/products/:id', async (req, res) => {
  try {
    const product = await prisma.product.findUnique({
      where: { id: parseInt(req.params.id) },
    });
    if (!product) {
      return res.status(404).json({ error: 'Product not found' });
    }
    res.json(product);
  } catch (error) {
    //res.status(500).json({ error: error.message });
    next(error); // Passa o erro para o middleware de tratamento de erros

  }
});

// Rota para atualizar um produto pelo ID
app.put('/api/products/:id', async (req, res) => {
  try {
    const product = await prisma.product.update({
      where: { id: parseInt(req.params.id) },
      data: req.body,
    });
    res.json(product);
  } catch (error) {
    //res.status(400).json({ error: error.message });
    next(error); // Passa o erro para o middleware de tratamento de erros

  }
});

// Rota para deletar um produto pelo ID
app.delete('/api/products/:id', async (req, res) => {
  try {
    await prisma.product.delete({
      where: { id: parseInt(req.params.id) },
    });
    res.json({ message: 'Product deleted' });
  } catch (error) {
    //res.status(500).json({ error: error.message });
    next(error); // Passa o erro para o middleware de tratamento de erros

  }
});


// Middleware de tratamento de erros
app.use((err, req, res, next) => {
  console.error(err.stack); // Loga os detalhes do erro no servidor
  res.status(500).json({ error: 'Internal server error' }); // Mensagem genérica para o cliente
});
5656

// Rota para servir o index.html do Angular
app.get('*', (req, res) => {
  console.log("June 06");
    res.sendFile(path.join(__dirname, FrontEnd, 'index.html'));
});
// Iniciar servidor
const PORT = process.env.PORT || 8080;
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});





import {
  Controller, Get, Post, Body, Param, Delete, Patch, ParseIntPipe, Headers,
} from '@nestjs/common';
import { ApiBearerAuth, ApiOperation, ApiParam, ApiResponse, ApiTags } from '@nestjs/swagger';
import { BooksService } from './books.service';
import { CreateBookDto } from './dto/create-book.dto';
import { UpdateBookDto } from './dto/update-book.dto';

@ApiTags('books')
@Controller('books')
export class BooksController {
  constructor(private readonly booksService: BooksService) {}

  @ApiOperation({ summary: 'Cria um livro' })
  @ApiResponse({ status: 201, description: 'Livro criado' })
  @ApiBearerAuth('jwt')
  @Post()
  create(@Body() dto: CreateBookDto, @Headers('authorization') _auth?: string) {
    // Para fins didáticos, não validaremos de fato o JWT aqui.
    return this.booksService.create(dto);
  }

  @ApiOperation({ summary: 'Lista todos os livros' })
  @Get()
  findAll() {
    return this.booksService.findAll();
  }

  @ApiOperation({ summary: 'Busca um livro pelo id' })
  @ApiParam({ name: 'id', type: Number, example: 1 })
  @Get(':id')
  findOne(@Param('id', ParseIntPipe) id: number) {
    return this.booksService.findOne(id);
  }

  @ApiOperation({ summary: 'Atualiza um livro' })
  @ApiParam({ name: 'id', type: Number, example: 1 })
  @Patch(':id')
  update(@Param('id', ParseIntPipe) id: number, @Body() dto: UpdateBookDto) {
    return this.booksService.update(id, dto);
  }

  @ApiOperation({ summary: 'Remove um livro' })
  @ApiParam({ name: 'id', type: Number, example: 1 })
  @ApiResponse({ status: 204, description: 'Removido' })
  @Delete(':id')
  remove(@Param('id', ParseIntPipe) id: number) {
    this.booksService.remove(id);
    return { status: 'ok' };
  }
}


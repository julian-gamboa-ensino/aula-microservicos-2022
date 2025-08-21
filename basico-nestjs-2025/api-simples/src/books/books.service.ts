import { Injectable, NotFoundException } from '@nestjs/common';
import { Book } from './entities/book.entity';
import { CreateBookDto } from './dto/create-book.dto';
import { UpdateBookDto } from './dto/update-book.dto';

@Injectable()
export class BooksService {
  private books: Book[] = [];
  private seq = 1;

  create(dto: CreateBookDto): Book {
    const book: Book = { id: this.seq++, ...dto };
    this.books.push(book);
    return book;
  }

  findAll(): Book[] {
    return this.books;
  }

  findOne(id: number): Book {
    const b = this.books.find((x) => x.id === id);
    if (!b) throw new NotFoundException(`Book ${id} não encontrado`);
    return b;
  }

  update(id: number, dto: UpdateBookDto): Book {
    const b = this.findOne(id);
    Object.assign(b, dto);
    return b;
  }

  remove(id: number): void {
    const idx = this.books.findIndex((x) => x.id === id);
    if (idx === -1) throw new NotFoundException(`Book ${id} não encontrado`);
    this.books.splice(idx, 1);
  }
}


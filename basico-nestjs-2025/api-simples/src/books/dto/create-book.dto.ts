import { ApiProperty } from '@nestjs/swagger';
import { IsString, IsInt, Min, IsOptional } from 'class-validator';

export class CreateBookDto {
  @ApiProperty({ example: 'Clean Code' })
  @IsString()
  title: string;

  @ApiProperty({ example: 'Robert C. Martin' })
  @IsString()
  author: string;

  @ApiProperty({ example: 2008, minimum: 0 })
  @IsInt()
  @Min(0)
  year: number;

  @ApiProperty({ example: 'Práticas de código limpo', required: false })
  @IsOptional()
  @IsString()
  description?: string;
}

import { Injectable, NotFoundException } from '@nestjs/common';
import { tours } from './data';
import { CreateTourDto } from './dto/create-tour.dto';
import { UpdateTourDto } from './dto/update-tour.dto';

@Injectable()
export class ToursService {
  create(createTourDto: CreateTourDto) {
    return 'This action adds a new tour';
  }

  findAll() {
    return tours;
  }

  findOne(id: string) {
    const tour = tours.find(({ _id }) => _id === id);

    if (!tour) {
      throw new NotFoundException('Tour not found');
    }

    return tour;
  }

  update(id: string, updateTourDto: UpdateTourDto) {
    const tour = tours.find(({ _id }) => _id === id);

    if (!tour) {
      throw new NotFoundException('Tour not found');
    }

    return tour;
  }

  remove(id: string) {
    const idx = tours.findIndex(({ _id }) => _id === id);

    if (idx === -1) {
      throw new NotFoundException('Tour not found');
    }

    tours.splice(idx, 1);
  }
}

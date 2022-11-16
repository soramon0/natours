import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { ToursModule } from './tours/tours.module';

@Module({
  imports: [
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    ToursModule,
  ],
})
export class AppModule {}

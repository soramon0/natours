import { VersioningType } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import {
  FastifyAdapter,
  NestFastifyApplication,
} from '@nestjs/platform-fastify';
import { AppModule } from './app.module';

const PORT = process.env.API_PORT || 4000;
const HOST = process.env.API_HOST || '0.0.0.0';

async function bootstrap() {
  const app = await NestFactory.create<NestFastifyApplication>(
    AppModule,
    new FastifyAdapter({ logger: true }),
  );

  app.enableVersioning({
    type: VersioningType.URI,
  });

  await app.listen(PORT, HOST);
}

bootstrap();

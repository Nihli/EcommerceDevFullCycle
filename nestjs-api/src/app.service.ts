import { Injectable } from '@nestjs/common';

//todo serviço precisa ter o decorator Injectable
@Injectable()
export class AppService {
  getHello(): string {
    return 'Hello World!';
  }
}

import * as cdk from 'aws-cdk-lib';
import { LambdaIntegration, RestApi } from 'aws-cdk-lib/aws-apigateway';
import { AttributeType, BillingMode, Table } from 'aws-cdk-lib/aws-dynamodb';
import { Runtime } from 'aws-cdk-lib/aws-lambda';
import { LogLevel, NodejsFunction } from 'aws-cdk-lib/aws-lambda-nodejs';
import { RetentionDays } from 'aws-cdk-lib/aws-logs';
import { Construct } from 'constructs';
import { join } from 'path';

export class ServerlessStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const natoursTable = new Table(this, 'NatoursTable', {
      billingMode: BillingMode.PROVISIONED,
      readCapacity: 5,
      writeCapacity: 5,
      partitionKey: { name: 'id', type: AttributeType.STRING },
      tableName: 'NatoursTable',
    });

    const helloFn = new NodejsFunction(this, 'HelloFunction', {
      bundling: {
        target: 'es2018',
        keepNames: true,
        logLevel: LogLevel.INFO,
        sourceMap: true,
        minify: true,
      },
      runtime: Runtime.NODEJS_16_X,
      logRetention: RetentionDays.ONE_DAY,
      entry: join(__dirname, '..', 'lambda', 'hello.ts'),
      environment: {
        TABLE: natoursTable.tableName,
        NODE_OPTIONS: '--enable-source-maps',
      },
    });

    natoursTable.grantReadData(helloFn);

    const api = new RestApi(this, 'Natours');

    api.root
      .addResource('hello')
      .addMethod('GET', new LambdaIntegration(helloFn));
  }
}

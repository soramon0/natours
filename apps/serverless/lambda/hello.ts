import { Context, APIGatewayProxyResult, APIGatewayEvent } from 'aws-lambda';
import {
  DynamoDBClient,
  GetItemCommand,
  GetItemCommandInput,
} from '@aws-sdk/client-dynamodb';
import { DynamoDBDocumentClient } from '@aws-sdk/lib-dynamodb'; // ES6 import

// eslint-disable-next-line turbo/no-undeclared-env-vars
const TABLE_NAME = process.env.TABLE;

if (!TABLE_NAME) {
  throw new Error(
    `Got "${TABLE_NAME}" for TABLE env variable, please define it.`
  );
}

const client = DynamoDBDocumentClient.from(
  new DynamoDBClient({ region: 'eu-west-3' })
);

export async function handler(
  event: APIGatewayEvent,
  context: Context
): Promise<APIGatewayProxyResult> {
  console.log('Start Hello function');
  console.log(`Handler working with "${TABLE_NAME}" table.`);
  console.log(`Event: ${JSON.stringify(event, null, 2)}`);
  console.log(`Context: ${JSON.stringify(context, null, 2)}`);

  const params: GetItemCommandInput = {
    TableName: TABLE_NAME,
    Key: {
      id: {
        S: '1',
      },
    },
    ProjectionExpression: 'ATTRIBUTE_NAME',
  };

  try {
    const data = await client.send(new GetItemCommand(params));
    console.log(`DB Response: ${JSON.stringify(data, null, 2)}`);
  } catch (error) {
    console.log(`DB GET Failed ${error}`);
  }

  return {
    statusCode: 200,
    headers: { 'Content-Type': 'text/json' },
    body: JSON.stringify({
      message: `Hello, CDK friends! You've hit ${event.path}`,
    }),
  };
}

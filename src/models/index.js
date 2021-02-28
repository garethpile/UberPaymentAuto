// @ts-check
import { initSchema } from '@aws-amplify/datastore';
import { schema } from './schema';



const { DriverStatements, UberDriver, UberTransactions, DriverTransactions } = initSchema(schema);

export {
  DriverStatements,
  UberDriver,
  UberTransactions,
  DriverTransactions
};
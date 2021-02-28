import { ModelInit, MutableModel, PersistentModelConstructor } from "@aws-amplify/datastore";





export declare class DriverStatements {
  readonly id: string;
  readonly uberdriverID?: string;
  readonly creationDate?: string;
  readonly documentLink?: string;
  constructor(init: ModelInit<DriverStatements>);
  static copyOf(source: DriverStatements, mutator: (draft: MutableModel<DriverStatements>) => MutableModel<DriverStatements> | void): DriverStatements;
}

export declare class UberDriver {
  readonly id: string;
  readonly firstName?: string;
  readonly lastName?: string;
  readonly uberId?: string;
  readonly UberTransactions?: (UberTransactions | null)[];
  readonly DriverTransactions?: (DriverTransactions | null)[];
  readonly DriverStatements?: (DriverStatements | null)[];
  constructor(init: ModelInit<UberDriver>);
  static copyOf(source: UberDriver, mutator: (draft: MutableModel<UberDriver>) => MutableModel<UberDriver> | void): UberDriver;
}

export declare class UberTransactions {
  readonly id: string;
  readonly total?: number;
  readonly trips?: number;
  readonly questPromotion?: number;
  readonly fare?: number;
  readonly waitTime?: number;
  readonly tip?: number;
  readonly toll?: number;
  readonly cashCollection?: number;
  readonly uberFee?: number;
  readonly cancellation?: number;
  readonly payouts?: number;
  readonly adjustedFare?: number;
  readonly date?: string;
  readonly uberdriverID?: string;
  constructor(init: ModelInit<UberTransactions>);
  static copyOf(source: UberTransactions, mutator: (draft: MutableModel<UberTransactions>) => MutableModel<UberTransactions> | void): UberTransactions;
}

export declare class DriverTransactions {
  readonly id: string;
  readonly date?: string;
  readonly description?: string;
  readonly amount?: number;
  readonly reviewed?: number;
  readonly uberdriverID?: string;
  constructor(init: ModelInit<DriverTransactions>);
  static copyOf(source: DriverTransactions, mutator: (draft: MutableModel<DriverTransactions>) => MutableModel<DriverTransactions> | void): DriverTransactions;
}
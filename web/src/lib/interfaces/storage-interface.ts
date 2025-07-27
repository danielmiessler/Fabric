export interface StorageEntity {
  Name: string;
  Description?: string;
  Pattern?: string | object;
  Session?: string | object;
  Context?: string | object;
}

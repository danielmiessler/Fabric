import TransformDescriptor, { toDescriptor } from '../TransformDescriptor';
import type TransformContext from './TransformContext';
import type Item from '../Item';
import type ItemResult from '../ItemResult';

/**
 * Transforms the incoming schema to what the transformer produces.
 */
type SchemaTransformer = (incomingSchema: string[]) => string[];

export default abstract class ItemTransformer {
  readonly name: string;
  readonly description: string;
  readonly descriptor: TransformDescriptor;
  readonly schemaTransformer: SchemaTransformer;

  constructor(
    name: string,
    description: string,
    descriptorPartial: Partial<TransformDescriptor>,
    schemaTransformer: SchemaTransformer = (schema) => schema,
  ) {
    this.name = name;
    this.description = description;
    this.descriptor = toDescriptor(descriptorPartial);
    this.schemaTransformer = schemaTransformer;
  }

  abstract transform(context: TransformContext, inputItems: Item[]): ItemResult;
}

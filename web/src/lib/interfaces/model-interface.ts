export interface VendorModel {
  name: string;
  vendor: string;
}

export interface ModelsResponse {
  models: string[];
  vendors: Record<string, string[]>;
}


export interface ModelConfig {
  model: string;
  temperature: number;
  top_p: number;
  maxLength: number;
  frequency: number;
  presence: number;
}

export type Bindings = {
  rnapi: KVNamespace;
  GCP_SERVICE_ACCOUNT: string;
}

declare global {
  function getMiniflareBindings(): Bindings
}

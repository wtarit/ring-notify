import { useEffect, useState, useCallback } from "react";
import { listAPIKeys, createAPIKey, revokeAPIKey } from "~/lib/api";
import { ConfirmDialog } from "~/components/confirm-dialog";
import { CopyButton } from "~/components/copy-button";
import { EmptyState } from "~/components/empty-state";
import type { APIKeyResponse } from "~/lib/types";

export function meta() {
  return [{ title: "API Keys - Ring Notify" }];
}

export default function APIKeysPage() {
  const [apiKeys, setApiKeys] = useState<APIKeyResponse[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Create modal state
  const [showCreate, setShowCreate] = useState(false);
  const [createName, setCreateName] = useState("");
  const [createExpiry, setCreateExpiry] = useState("");
  const [creating, setCreating] = useState(false);

  // Newly created key (show once)
  const [newKey, setNewKey] = useState<string | null>(null);

  // Revoke confirmation state
  const [revokingKey, setRevokingKey] = useState<APIKeyResponse | null>(null);

  const fetchKeys = useCallback(async () => {
    try {
      setError(null);
      const data = await listAPIKeys();
      setApiKeys(data.apiKeys || []);
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchKeys();
  }, [fetchKeys]);

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();
    setCreating(true);
    try {
      const payload: { name: string; expiresAt?: string } = { name: createName };
      if (createExpiry) {
        payload.expiresAt = new Date(createExpiry).toISOString();
      }
      const created = await createAPIKey(payload);
      setNewKey(created.key);
      setShowCreate(false);
      setCreateName("");
      setCreateExpiry("");
      fetchKeys();
    } catch (err: any) {
      setError(err.message);
    } finally {
      setCreating(false);
    }
  };

  const handleRevoke = async () => {
    if (!revokingKey) return;
    try {
      await revokeAPIKey(revokingKey.id);
      setRevokingKey(null);
      fetchKeys();
    } catch (err: any) {
      setError(err.message);
    }
  };

  if (loading) {
    return (
      <div className="flex justify-center py-16">
        <span className="loading loading-spinner loading-lg" />
      </div>
    );
  }

  return (
    <div>
      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-bold">API Keys</h1>
        <button className="btn btn-primary btn-sm" onClick={() => setShowCreate(true)}>
          Create API Key
        </button>
      </div>

      {error && (
        <div className="alert alert-error mb-4">
          <span>{error}</span>
        </div>
      )}

      {/* Newly created key alert */}
      {newKey && (
        <div className="alert alert-warning mb-4">
          <div className="flex-1">
            <p className="font-bold">Copy your API key now. It won't be shown again.</p>
            <code className="block mt-2 break-all">{newKey}</code>
          </div>
          <div className="flex gap-2">
            <CopyButton text={newKey} label="Copy" />
            <button className="btn btn-sm btn-ghost" onClick={() => setNewKey(null)}>
              Dismiss
            </button>
          </div>
        </div>
      )}

      {/* API Key list */}
      {apiKeys.length === 0 ? (
        <EmptyState
          title="No API keys"
          description="Create an API key to start sending notifications from your devices and services."
          action={
            <button className="btn btn-primary btn-sm" onClick={() => setShowCreate(true)}>
              Create API Key
            </button>
          }
        />
      ) : (
        <>
          {/* Desktop table */}
          <div className="hidden md:block overflow-x-auto">
            <table className="table">
              <thead>
                <tr>
                  <th>Name</th>
                  <th>Key</th>
                  <th>Created</th>
                  <th>Expires</th>
                  <th>Last Used</th>
                  <th>Status</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {apiKeys.map((key) => (
                  <tr key={key.id}>
                    <td className="font-medium">{key.name}</td>
                    <td>
                      <code className="text-sm">{key.key}</code>
                      <CopyButton text={key.key} className="ml-1" />
                    </td>
                    <td>{new Date(key.createdAt).toLocaleDateString()}</td>
                    <td>{key.expiresAt ? new Date(key.expiresAt).toLocaleDateString() : "Never"}</td>
                    <td>{key.lastUsedAt ? new Date(key.lastUsedAt).toLocaleDateString() : "Never"}</td>
                    <td>
                      <span className={`badge ${key.isActive ? "badge-success" : "badge-error"}`}>
                        {key.isActive ? "Active" : "Revoked"}
                      </span>
                    </td>
                    <td>
                      {key.isActive && (
                        <button
                          className="btn btn-sm btn-ghost text-error"
                          onClick={() => setRevokingKey(key)}
                        >
                          Revoke
                        </button>
                      )}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>

          {/* Mobile card list */}
          <div className="md:hidden space-y-3">
            {apiKeys.map((key) => (
              <div key={key.id} className="card bg-base-200">
                <div className="card-body p-4">
                  <div className="flex items-center justify-between">
                    <h3 className="font-medium">{key.name}</h3>
                    <span className={`badge badge-sm ${key.isActive ? "badge-success" : "badge-error"}`}>
                      {key.isActive ? "Active" : "Revoked"}
                    </span>
                  </div>
                  <div className="flex items-center gap-2 mt-1">
                    <code className="text-sm break-all">{key.key}</code>
                    <CopyButton text={key.key} />
                  </div>
                  <div className="text-sm opacity-70 space-y-1 mt-2">
                    <p>Created: {new Date(key.createdAt).toLocaleDateString()}</p>
                    <p>Expires: {key.expiresAt ? new Date(key.expiresAt).toLocaleDateString() : "Never"}</p>
                    <p>Last Used: {key.lastUsedAt ? new Date(key.lastUsedAt).toLocaleDateString() : "Never"}</p>
                  </div>
                  {key.isActive && (
                    <div className="card-actions justify-end mt-2">
                      <button
                        className="btn btn-sm btn-ghost text-error"
                        onClick={() => setRevokingKey(key)}
                      >
                        Revoke
                      </button>
                    </div>
                  )}
                </div>
              </div>
            ))}
          </div>
        </>
      )}

      {/* Create modal */}
      {showCreate && (
        <dialog className="modal modal-open">
          <div className="modal-box">
            <h3 className="font-bold text-lg">Create API Key</h3>
            <form onSubmit={handleCreate} className="mt-4 space-y-3">
              <label className="floating-label">
                <span>Name</span>
                <input
                  type="text"
                  placeholder="Name"
                  className="input input-bordered w-full"
                  value={createName}
                  onChange={(e) => setCreateName(e.target.value)}
                  required
                />
              </label>
              <label className="floating-label">
                <span>Expiry Date (optional)</span>
                <input
                  type="date"
                  placeholder="Expiry Date (optional)"
                  className="input input-bordered w-full"
                  value={createExpiry}
                  onChange={(e) => setCreateExpiry(e.target.value)}
                  min={new Date().toISOString().split("T")[0]}
                />
              </label>
              <div className="modal-action">
                <button type="button" className="btn" onClick={() => setShowCreate(false)}>
                  Cancel
                </button>
                <button type="submit" className="btn btn-primary" disabled={creating || !createName.trim()}>
                  {creating ? <span className="loading loading-spinner loading-sm" /> : "Create"}
                </button>
              </div>
            </form>
          </div>
          <form method="dialog" className="modal-backdrop">
            <button onClick={() => setShowCreate(false)}>close</button>
          </form>
        </dialog>
      )}

      {/* Revoke confirmation */}
      <ConfirmDialog
        open={!!revokingKey}
        title="Revoke API Key"
        message={`Are you sure you want to revoke "${revokingKey?.name}"? Services using this key will stop working.`}
        confirmLabel="Revoke"
        onConfirm={handleRevoke}
        onCancel={() => setRevokingKey(null)}
      />
    </div>
  );
}

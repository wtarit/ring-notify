import { useEffect, useState, useCallback } from "react";
import { QRCodeSVG } from "qrcode.react";
import { listDevices, updateDevice, deleteDevice } from "~/lib/api";
import { isWebView } from "~/lib/webview";
import { ConfirmDialog } from "~/components/confirm-dialog";
import { EmptyState } from "~/components/empty-state";
import type { DeviceResponse } from "~/lib/types";

const GITHUB_RELEASES_URL = "https://github.com/nicepkg/ring-notify/releases";

export function meta() {
  return [{ title: "Devices - Ring Notify" }];
}

export default function DevicesPage() {
  const [devices, setDevices] = useState<DeviceResponse[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Edit modal state
  const [editingDevice, setEditingDevice] = useState<DeviceResponse | null>(null);
  const [editName, setEditName] = useState("");
  const [editSaving, setEditSaving] = useState(false);

  // Delete confirmation state
  const [deletingDevice, setDeletingDevice] = useState<DeviceResponse | null>(null);

  const inWebView = isWebView();

  const fetchDevices = useCallback(async () => {
    try {
      setError(null);
      const data = await listDevices();
      setDevices(data.devices || []);
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchDevices();
  }, [fetchDevices]);

  const handleEdit = (device: DeviceResponse) => {
    setEditingDevice(device);
    setEditName(device.deviceName);
  };

  const handleEditSave = async () => {
    if (!editingDevice) return;
    setEditSaving(true);
    try {
      await updateDevice(editingDevice.id, { deviceName: editName });
      setEditingDevice(null);
      fetchDevices();
    } catch (err: any) {
      setError(err.message);
    } finally {
      setEditSaving(false);
    }
  };

  const handleDelete = async () => {
    if (!deletingDevice) return;
    try {
      await deleteDevice(deletingDevice.id);
      setDeletingDevice(null);
      fetchDevices();
    } catch (err: any) {
      setError(err.message);
    }
  };

  const handleRegisterWebView = () => {
    // Request FCM token from the native React Native app
    (window as any).ReactNativeWebView?.postMessage(
      JSON.stringify({ type: "REQUEST_FCM_TOKEN" })
    );
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
        <h1 className="text-2xl font-bold">Devices</h1>
      </div>

      {error && (
        <div className="alert alert-error mb-4">
          <span>{error}</span>
        </div>
      )}

      {/* WebView: Register this device */}
      {inWebView && (
        <div className="card bg-base-200 mb-6">
          <div className="card-body">
            <h2 className="card-title">Register This Device</h2>
            <p className="opacity-70">Add this device to receive call notifications.</p>
            <div className="card-actions justify-end">
              <button className="btn btn-primary" onClick={handleRegisterWebView}>
                Register This Device
              </button>
            </div>
          </div>
        </div>
      )}

      {/* Browser: QR code to install app */}
      {!inWebView && (
        <div className="card bg-base-200 mb-6">
          <div className="card-body items-center text-center">
            <h2 className="card-title">Install the Mobile App</h2>
            <p className="opacity-70 mb-4">
              Scan the QR code or use the link below to download the Ring Notify app.
            </p>
            <QRCodeSVG value={GITHUB_RELEASES_URL} size={160} />
            <a
              href={GITHUB_RELEASES_URL}
              target="_blank"
              rel="noopener noreferrer"
              className="link link-primary mt-2 text-sm"
            >
              {GITHUB_RELEASES_URL}
            </a>
          </div>
        </div>
      )}

      {/* Device list */}
      {devices.length === 0 ? (
        <EmptyState
          title="No devices registered"
          description="Install the mobile app and register your device to start receiving call notifications."
        />
      ) : (
        <>
          {/* Desktop table */}
          <div className="hidden md:block overflow-x-auto">
            <table className="table">
              <thead>
                <tr>
                  <th>Name</th>
                  <th>Type</th>
                  <th>Registered</th>
                  <th>Last Active</th>
                  <th>Status</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {devices.map((device) => (
                  <tr key={device.id}>
                    <td className="font-medium">{device.deviceName}</td>
                    <td className="capitalize">{device.deviceType}</td>
                    <td>{new Date(device.registeredAt).toLocaleDateString()}</td>
                    <td>{new Date(device.lastActive).toLocaleDateString()}</td>
                    <td>
                      <span className={`badge ${device.isActive ? "badge-success" : "badge-error"}`}>
                        {device.isActive ? "Active" : "Inactive"}
                      </span>
                    </td>
                    <td className="space-x-1">
                      <button className="btn btn-sm btn-ghost" onClick={() => handleEdit(device)}>
                        Edit
                      </button>
                      <button
                        className="btn btn-sm btn-ghost text-error"
                        onClick={() => setDeletingDevice(device)}
                      >
                        Delete
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>

          {/* Mobile card list */}
          <div className="md:hidden space-y-3">
            {devices.map((device) => (
              <div key={device.id} className="card bg-base-200">
                <div className="card-body p-4">
                  <div className="flex items-center justify-between">
                    <h3 className="font-medium">{device.deviceName}</h3>
                    <span className={`badge badge-sm ${device.isActive ? "badge-success" : "badge-error"}`}>
                      {device.isActive ? "Active" : "Inactive"}
                    </span>
                  </div>
                  <div className="text-sm opacity-70 space-y-1">
                    <p>Type: <span className="capitalize">{device.deviceType}</span></p>
                    <p>Registered: {new Date(device.registeredAt).toLocaleDateString()}</p>
                    <p>Last Active: {new Date(device.lastActive).toLocaleDateString()}</p>
                  </div>
                  <div className="card-actions justify-end mt-2">
                    <button className="btn btn-sm btn-ghost" onClick={() => handleEdit(device)}>
                      Edit
                    </button>
                    <button
                      className="btn btn-sm btn-ghost text-error"
                      onClick={() => setDeletingDevice(device)}
                    >
                      Delete
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </>
      )}

      {/* Edit modal */}
      {editingDevice && (
        <dialog className="modal modal-open">
          <div className="modal-box">
            <h3 className="font-bold text-lg">Rename Device</h3>
            <div className="mt-4">
              <label className="floating-label">
                <span>Device Name</span>
                <input
                  type="text"
                  placeholder="Device Name"
                  className="input input-bordered w-full"
                  value={editName}
                  onChange={(e) => setEditName(e.target.value)}
                />
              </label>
            </div>
            <div className="modal-action">
              <button className="btn" onClick={() => setEditingDevice(null)}>
                Cancel
              </button>
              <button
                className="btn btn-primary"
                onClick={handleEditSave}
                disabled={editSaving || !editName.trim()}
              >
                {editSaving ? <span className="loading loading-spinner loading-sm" /> : "Save"}
              </button>
            </div>
          </div>
          <form method="dialog" className="modal-backdrop">
            <button onClick={() => setEditingDevice(null)}>close</button>
          </form>
        </dialog>
      )}

      {/* Delete confirmation */}
      <ConfirmDialog
        open={!!deletingDevice}
        title="Delete Device"
        message={`Are you sure you want to remove "${deletingDevice?.deviceName}"? This device will no longer receive notifications.`}
        onConfirm={handleDelete}
        onCancel={() => setDeletingDevice(null)}
      />
    </div>
  );
}

import { useState } from "react";
import { Link, Navigate, useNavigate } from "react-router";
import { useAuth } from "~/context/auth-context";
import { supabase } from "~/lib/supabase";

export function meta() {
  return [{ title: "Set New Password - Ring Notify" }];
}

export default function ResetPassword() {
  const { session, loading } = useAuth();
  const navigate = useNavigate();
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [submitting, setSubmitting] = useState(false);

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <span className="loading loading-spinner loading-lg" />
      </div>
    );
  }

  // User must be signed in (via the recovery link) to reset password
  if (!session) {
    return <Navigate to="/login" replace />;
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSubmitting(true);
    try {
      if (password.length < 8) {
        setError("Password must be at least 8 characters.");
        setSubmitting(false);
        return;
      }
      if (password !== confirmPassword) {
        setError("Passwords do not match.");
        setSubmitting(false);
        return;
      }
      const { error } = await supabase.auth.updateUser({ password });
      if (error) throw error;
      navigate("/dashboard/devices", { replace: true });
    } catch (err: any) {
      setError(err.message);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <div className="card w-full max-w-sm bg-base-200 shadow-xl">
        <div className="card-body">
          <h2 className="card-title justify-center text-2xl">Ring Notify</h2>
          <p className="text-center opacity-70">Set your new password</p>

          {error && (
            <div className="alert alert-error mt-2">
              <span>{error}</span>
            </div>
          )}

          <form onSubmit={handleSubmit} className="mt-4 space-y-3">
            <label className="floating-label">
              <span>New Password</span>
              <input
                type="password"
                placeholder="New Password"
                className="input input-bordered w-full"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                minLength={8}
              />
            </label>
            <label className="floating-label">
              <span>Confirm Password</span>
              <input
                type="password"
                placeholder="Confirm Password"
                className="input input-bordered w-full"
                value={confirmPassword}
                onChange={(e) => setConfirmPassword(e.target.value)}
                required
                minLength={8}
              />
            </label>
            <button type="submit" className="btn btn-primary w-full" disabled={submitting}>
              {submitting ? (
                <span className="loading loading-spinner loading-sm" />
              ) : (
                "Update Password"
              )}
            </button>
          </form>

          <div className="text-center mt-4">
            <Link to="/login" className="link link-hover text-sm opacity-70">
              Back to Sign In
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}

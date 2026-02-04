import { useState } from "react";
import { Link, Navigate } from "react-router";
import { useAuth } from "~/context/auth-context";
import { supabase } from "~/lib/supabase";

export function meta() {
  return [{ title: "Reset Password - Ring Notify" }];
}

export default function ForgotPassword() {
  const { session, loading } = useAuth();
  const [email, setEmail] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [submitting, setSubmitting] = useState(false);
  const [resetSent, setResetSent] = useState(false);

  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <span className="loading loading-spinner loading-lg" />
      </div>
    );
  }

  if (session) {
    return <Navigate to="/dashboard/devices" replace />;
  }

  const handleSubmit = async (e: React.SubmitEvent) => {
    e.preventDefault();
    setError(null);
    setSubmitting(true);
    try {
      const { error } = await supabase.auth.resetPasswordForEmail(email, {
        redirectTo: `${window.location.origin}/auth/callback`,
      });
      if (error) throw error;
      setResetSent(true);
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
          <p className="text-center opacity-70">Reset your password</p>

          {error && (
            <div className="alert alert-error mt-2">
              <span>{error}</span>
            </div>
          )}

          {resetSent ? (
            <div className="mt-4 text-center space-y-4">
              <div className="alert alert-success">
                <span>Check your email for a password reset link.</span>
              </div>
              <Link to="/login" className="btn btn-ghost btn-sm">
                Back to Sign In
              </Link>
            </div>
          ) : (
            <form onSubmit={handleSubmit} className="mt-4 space-y-3">
              <label className="floating-label">
                <span>Email</span>
                <input
                  type="email"
                  placeholder="Email"
                  className="input input-bordered w-full"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  required
                />
              </label>
              <button type="submit" className="btn btn-primary w-full" disabled={submitting}>
                {submitting ? (
                  <span className="loading loading-spinner loading-sm" />
                ) : (
                  "Send Reset Link"
                )}
              </button>
            </form>
          )}

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

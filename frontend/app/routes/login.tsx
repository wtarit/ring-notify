import { useState } from "react";
import { FcGoogle } from "react-icons/fc";
import { Link, Navigate } from "react-router";
import { useAuth } from "~/context/auth-context";
import { supabase } from "~/lib/supabase";

export function meta() {
  return [{ title: "Sign In - Ring Notify" }];
}

export default function Login() {
  const { session, loading } = useAuth();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [submitting, setSubmitting] = useState(false);

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

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSubmitting(true);
    try {
      const { error } = await supabase.auth.signInWithPassword({ email, password });
      if (error) throw error;
    } catch (err: any) {
      setError(err.message);
    } finally {
      setSubmitting(false);
    }
  };

  const handleGoogleAuth = async () => {
    const { error } = await supabase.auth.signInWithOAuth({
      provider: "google",
      options: { redirectTo: `${window.location.origin}/auth/callback` },
    });
    if (error) setError(error.message);
  };

  return (
    <div className="min-h-screen flex items-center justify-center p-4">
      <div className="card w-full max-w-sm bg-base-200 shadow-xl">
        <div className="card-body">
          <h2 className="card-title justify-center text-2xl">Ring Notify</h2>
          <p className="text-center opacity-70">Sign in to your account</p>

          {error && (
            <div className="alert alert-error mt-2">
              <span>{error}</span>
            </div>
          )}

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
            <label className="floating-label">
              <span>Password</span>
              <input
                type="password"
                placeholder="Password"
                className="input input-bordered w-full"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                minLength={8}
              />
            </label>
            <div className="text-right">
              <Link to="/forgot-password" className="link link-hover text-sm opacity-70">
                Forgot password?
              </Link>
            </div>
            <button type="submit" className="btn btn-primary w-full" disabled={submitting}>
              {submitting ? <span className="loading loading-spinner loading-sm" /> : "Sign In"}
            </button>
          </form>

          <div className="divider">or</div>

          <button className="btn btn-outline w-full" onClick={handleGoogleAuth}>
            <FcGoogle />
            Continue with Google
          </button>

          <p className="text-center mt-4 text-sm">
            Don't have an account?{" "}
            <Link to="/signup" className="link link-primary">
              Sign Up
            </Link>
          </p>

          <div className="text-center mt-2">
            <Link to="/" className="link link-hover text-sm opacity-70">
              Back to Home
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}

import { useState } from "react";
import { FcGoogle } from "react-icons/fc";
import { Link, Navigate } from "react-router";
import { useAuth } from "~/context/auth-context";
import { supabase } from "~/lib/supabase";

export function meta() {
  return [{ title: "Sign Up - Ring Notify" }];
}

export default function SignUp() {
  const { session, loading } = useAuth();
  const [email, setEmail] = useState("");
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

  if (session) {
    return <Navigate to="/dashboard/devices" replace />;
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
      const { error } = await supabase.auth.signUp({ email, password });
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
          <p className="text-center opacity-70">Create your account</p>

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
              {submitting ? <span className="loading loading-spinner loading-sm" /> : "Sign Up"}
            </button>
          </form>

          <div className="divider">or</div>

          <button className="btn btn-outline w-full" onClick={handleGoogleAuth}>
            <FcGoogle />
            Continue with Google
          </button>

          <p className="text-center mt-4 text-sm">
            Already have an account?{" "}
            <Link to="/login" className="link link-primary">
              Sign In
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

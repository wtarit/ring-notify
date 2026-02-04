import { useEffect } from "react";
import { useNavigate } from "react-router";
import { supabase } from "~/lib/supabase";

export function meta() {
  return [{ title: "Authenticating... - Ring Notify" }];
}

export default function AuthCallback() {
  const navigate = useNavigate();

  useEffect(() => {
    supabase.auth.onAuthStateChange((event) => {
      if (event === "PASSWORD_RECOVERY") {
        navigate("/reset-password", { replace: true });
      } else if (event === "SIGNED_IN") {
        navigate("/dashboard/devices", { replace: true });
      }
    });
  }, [navigate]);

  return (
    <div className="min-h-screen flex items-center justify-center">
      <span className="loading loading-spinner loading-lg" />
    </div>
  );
}

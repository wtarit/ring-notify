import { Outlet, Navigate, useLocation } from "react-router";
import { ProtectedRoute } from "~/components/protected-route";
import { Navbar } from "~/components/navbar";

export default function DashboardLayout() {
  const location = useLocation();

  // Redirect /dashboard to /dashboard/devices
  if (location.pathname === "/dashboard") {
    return <Navigate to="/dashboard/devices" replace />;
  }

  return (
    <ProtectedRoute>
      <div className="min-h-screen">
        <Navbar />
        <main className="container mx-auto p-4 max-w-5xl">
          <Outlet />
        </main>
      </div>
    </ProtectedRoute>
  );
}

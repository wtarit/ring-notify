import { Link } from "react-router";
import type { Route } from "./+types/home";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Ring Notify - When Notifications Aren't Enough" },
    { name: "description", content: "Turn any alert into an incoming call. Never miss what matters." },
  ];
}

const features = [
  {
    title: "Instant Call Alerts",
    description:
      "Convert any notification into a phone call that rings your device. No more missed alerts buried in your notification shade.",
  },
  {
    title: "REST API Integration",
    description:
      "Simple API to trigger calls from any device or service. Connect your IoT sensors, servers, or automation workflows.",
  },
  {
    title: "Multi-Device Support",
    description:
      "Register multiple devices and receive calls on all of them simultaneously. Stay reachable wherever you are.",
  },
];

export default function Home() {
  return (
    <div className="min-h-screen">
      {/* Hero */}
      <div className="hero min-h-[70vh]">
        <div className="hero-content text-center">
          <div className="max-w-2xl">
            <h1 className="text-5xl font-bold">When Notifications Aren't Enough</h1>
            <p className="py-6 text-lg opacity-80">
              Turn any alert into an incoming call. Never miss what matters.
            </p>
            <div className="flex flex-wrap justify-center gap-3">
              <Link to="/signup" className="btn btn-primary">
                Get Started
              </Link>
              <a
                href="https://ringnotify.wtarit.me"
                target="_blank"
                rel="noopener noreferrer"
                className="btn btn-outline"
              >
                Documentation
              </a>
            </div>
          </div>
        </div>
      </div>

      {/* Features */}
      <div className="container mx-auto px-4 pb-20 max-w-5xl">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          {features.map((feature) => (
            <div key={feature.title} className="card bg-base-200">
              <div className="card-body">
                <h2 className="card-title">{feature.title}</h2>
                <p className="opacity-70">{feature.description}</p>
              </div>
            </div>
          ))}
        </div>

        {/* App download */}
        <div className="text-center mt-16">
          <p className="opacity-70">
            Download the mobile app from{" "}
            <a
              href="https://github.com/nicepkg/ring-notify/releases"
              target="_blank"
              rel="noopener noreferrer"
              className="link link-primary"
            >
              GitHub Releases
            </a>
          </p>
        </div>
      </div>
    </div>
  );
}

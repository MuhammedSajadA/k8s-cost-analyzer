import { Routes, Route } from "react-router-dom";
import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";
import Cluster from "./pages/Cluster";
import Namespace from "./pages/Namespace";

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Login />} />
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="/clusters/:id" element={<Cluster />} />
      <Route path="/clusters/:id/namespaces/:ns" element={<Namespace />} />
    </Routes>
  );
}

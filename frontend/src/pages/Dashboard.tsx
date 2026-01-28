import { useEffect, useState } from "react";
import api from "../api/client";
import { useNavigate } from "react-router-dom";

export default function Dashboard() {
  const [clusters, setClusters] = useState<any[]>([]);
  const nav = useNavigate();

  useEffect(() => {
    api.get("/clusters").then(res => setClusters(res.data));
  }, []);

  return (
    <div className="p-6">
      <h1 className="text-2xl mb-4">Clusters</h1>

      {clusters.map(c => (
        <div
          key={c.id}
          className="border p-4 mb-2 cursor-pointer hover:bg-gray-50"
          onClick={() => nav(`/clusters/${c.id}`)}
        >
          {c.name}
        </div>
      ))}
    </div>
  );
}

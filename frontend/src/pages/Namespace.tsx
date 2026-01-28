import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import api from "../api/client";

export default function Namespace() {
  const { id, ns } = useParams();
  const [m, setM] = useState<any>(null);

  useEffect(() => {
    api.get(`/clusters/${id}/namespaces/${ns}/metrics`)
      .then(res => setM(res.data));
  }, [id, ns]);

  if (!m) return <div className="p-6">Loading...</div>;

  return (
    <div className="p-6">
      <h1 className="text-2xl mb-4">{m.namespace}</h1>

      <div className="grid grid-cols-3 gap-4">
        <Card title="CPU Requested" value={`${m.cpu_requested_m} m`} />
        <Card title="CPU Used" value={`${m.cpu_used_m} m`} />
        <Card title="CPU Waste" value={`${m.cpu_waste_m} m`} />

        <Card title="Memory Requested" value={`${(m.mem_requested_bytes/1024/1024).toFixed(0)} MB`} />
        <Card title="Memory Used" value={`${(m.mem_used_bytes/1024/1024).toFixed(0)} MB`} />
        <Card title="Memory Waste" value={`${(m.mem_waste_bytes/1024/1024).toFixed(0)} MB`} />
      </div>
    </div>
  );
}

function Card({ title, value }: any) {
  return (
    <div className="border p-4 rounded">
      <p className="text-gray-500">{title}</p>
      <p className="text-xl font-bold">{value}</p>
    </div>
  );
}

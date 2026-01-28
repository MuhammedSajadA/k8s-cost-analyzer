import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import api from "../api/client";

export default function Cluster() {
  const { id } = useParams();
  const [namespaces, setNamespaces] = useState<string[]>([]);
  const nav = useNavigate();

  useEffect(() => {
    api.get(`/clusters/${id}/namespaces`)
      .then(res => setNamespaces(res.data.namespaces));
  }, [id]);

  return (
    <div className="p-6">
      <h2 className="text-xl mb-4">Namespaces</h2>

      {namespaces.map(ns => (
        <div
          key={ns}
          className="border p-3 mb-2 cursor-pointer"
          onClick={() => nav(`/clusters/${id}/namespaces/${ns}`)}
        >
          {ns}
        </div>
      ))}
    </div>
  );
}

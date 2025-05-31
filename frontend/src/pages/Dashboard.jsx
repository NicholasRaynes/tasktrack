import { useEffect, useState } from "react";
import axios from "axios";

function Dashboard() {
  const [projects, setProjects] = useState([]);

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const token = localStorage.getItem("token");
        const res = await axios.get("http://localhost:8080/api/projects", {
          headers: { Authorization: `Bearer ${token}` },
        });
        setProjects(res.data);
      } catch (err) {
        console.error("Failed to fetch projects");
      }
    };
    fetchProjects();
  }, []);

  return (
    <div>
      <h2>Your Projects</h2>
      <ul>
        {projects.map((proj) => (
          <li key={proj.ID}>
            {proj.title}
            <ul>
              {proj.tasks?.map((task) => (
                <li key={task.ID}>{task.title}</li>
              ))}
            </ul>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default Dashboard;
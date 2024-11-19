// pages/index.js
import { useEffect, useState } from 'react';
import { fetchData } from './../../lib/api';
import Post from './../../components/post/post';

const HomePage = () => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const getData = async () => {
      try {
        const result = await fetchData();
        setData(result);
      } catch (error) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };

    getData();
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <div>
      <h1>Data from API</h1>
      {data.map(post => (
        <Post key={post.id} post={post} />
      ))}
    </div>
  );
};

export default HomePage;
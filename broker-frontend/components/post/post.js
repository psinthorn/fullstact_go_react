const Post = ({ post }) => {
  return (
    <div>
      const postTitle: string = post.title;
      const postBody: string = post.body;
      <h2>{post.title}</h2>
      <p>{post.body}</p>
    </div>
  );
};

export default Post;
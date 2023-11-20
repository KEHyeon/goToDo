import axios from "axios";
import * as S from "./todoInput.style.ts";
import { useState } from "react";

const TodoInput = () => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  const fetchPostTodo = async () => {
    try {
      if (title == "" || content == "") {
        alert("Fill out the form!");
        return;
      }
      await axios.post("http://localhost:8080/todo", { title, content });
      setTitle("");
      setContent("");
      alert("Success!");
      window.location.reload();
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <S.Contain>
      INPUT
      <S.Form>
        <div>
          Todo
          <S.Input
            value={title}
            onChange={(e) => {
              setTitle(e.target.value);
            }}
          />
        </div>
        <div>
          Description
          <S.Input
            value={content}
            onChange={(e) => {
              setContent(e.target.value);
            }}
          />
        </div>
        <S.Button onClick={fetchPostTodo}>GO!</S.Button>
      </S.Form>
    </S.Contain>
  );
};

export default TodoInput;

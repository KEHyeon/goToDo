import axios from "axios";
import * as S from "./index.style";
import { useEffect, useState } from "react";
import TodoItem from "./todoItem";

const TodoList = () => {
  const [todoList, setTodoList] = useState([]);

  const FetchTodoList = async () => {
    try {
      const res = await axios.get("http://localhost:8080/todo");
      setTodoList(res.data);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    FetchTodoList();
  }, []);
  return (
    <S.Contain>
      {todoList.map((data) => {
        console.log(data);
        return <TodoItem todo={data}></TodoItem>;
      })}
    </S.Contain>
  );
};

export default TodoList;

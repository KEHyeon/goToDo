import axios from "axios";
import * as S from "./index.style";
import { useEffect, useState } from "react";
import TodoItem from "./todoItem";
import Input from "./todoInput";

const TodoList = () => {
  const [todoList, setTodoList] = useState([]);
  const [visibleList, setVisibleTodoList] = useState([]);
  const [page, setPage] = useState(0);
  const FetchTodoList = async () => {
    try {
      const res = await axios.get("http://localhost:8080/todo");
      setTodoList(res.data);
    } catch (err) {
      console.error(err);
    }
  };
  const pagination = () => {
    setVisibleTodoList(
      todoList.filter((i, idx) => {
        return idx >= page * 5 && idx < page * 5 + 5;
      })
    );
  };

  useEffect(() => {
    FetchTodoList();
  }, []);
  useEffect(() => {
    pagination();
  }, [page, todoList]);
  return (
    <S.Contain>
      {visibleList.map((data) => {
        return <TodoItem todo={data}></TodoItem>;
      })}
      <S.ButtonWrapper>
        <div
          onClick={() => {
            if (page >= 1) setPage(page - 1);
          }}
        >
          pre
        </div>
        <div
          onClick={() => {
            if ((page + 1) * 5 < todoList.length) setPage(page + 1);
          }}
        >
          next
        </div>
      </S.ButtonWrapper>
      <Input setTodoList={setTodoList} todoList={todoList} />
    </S.Contain>
  );
};

export default TodoList;

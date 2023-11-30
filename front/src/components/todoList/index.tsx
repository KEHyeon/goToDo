import axios from "axios";
import * as S from "./index.style";
import { useEffect, useState } from "react";
import TodoItem from "./todoItem";
import Input from "./todoInput";
import { DragDropContext, Draggable, Droppable } from "react-beautiful-dnd";

const TodoList = () => {
  const [todoList, setTodoList] = useState([]);
  const [visibleList, setVisibleTodoList] = useState([]);
  const [page, setPage] = useState(0);
  const MAX = 5;
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
        return idx >= page * MAX && idx < page * MAX + MAX;
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
      <DragDropContext
        onDragEnd={(result) => {
          if (!result.destination) {
            return;
          }
          console.log(result);
          const updatedList = [...visibleList];
          const [removed] = updatedList.splice(result.source.index, 1);
          updatedList.splice(result.destination.index, 0, removed);
          setVisibleTodoList(updatedList);
        }}
      >
        <Droppable droppableId="Todos">
          {(provided) => (
            <S.DraggableWrapper
              {...provided.droppableProps}
              ref={provided.innerRef}
            >
              {visibleList.map((data: any, i) => (
                <Draggable
                  key={data.id}
                  draggableId={data.id.toString()}
                  index={i}
                >
                  {(provided) => (
                    <TodoItem todo={data} provided={provided}></TodoItem>
                  )}
                </Draggable>
              ))}
              {provided.placeholder}
            </S.DraggableWrapper>
          )}
        </Droppable>
      </DragDropContext>
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
            if ((page + 1) * MAX < todoList.length) setPage(page + 1);
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

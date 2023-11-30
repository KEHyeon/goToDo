import axios from "axios";
import * as S from "./index.style";
import { useEffect, useState } from "react";
import TodoItem from "./todoItem";
import Input from "./todoInput";
import { DragDropContext, Draggable, Droppable } from "react-beautiful-dnd";

const TodoList = () => {
  const [todoList, setTodoList] = useState([]);
  const [curDate, setCurDate] = useState(new Date());
  const MAX = 5;

  const formatDate = (date: Date) => {
    const formattedDate = `${date.getFullYear()}-${(date.getMonth() + 1)
      .toString()
      .padStart(2, "0")}-${date.getDate().toString().padStart(2, "0")}`;
    return formattedDate;
  };
  const FetchTodoList = async () => {
    try {
      const res = await axios.get(
        "http://localhost:8080/todo/" + formatDate(curDate)
      );
      setTodoList(res.data);
    } catch (err) {
      console.error(err);
    }
  };
  useEffect(() => {
    FetchTodoList();
  }, [curDate]);
  return (
    <S.Contain>
      <S.DateWrapper>
        <S.DateButtonWrapper>
          <div
            onClick={() => {
              const temp = new Date(curDate);
              temp.setDate(temp.getDate() - 1);
              setCurDate(temp);
              console.log(temp);
            }}
          >
            before
          </div>
          <div
            onClick={() => {
              const temp = new Date(curDate);
              temp.setDate(temp.getDate() + 1);
              setCurDate(temp);
              console.log(temp);
            }}
          >
            next
          </div>
        </S.DateButtonWrapper>
        <span>{formatDate(curDate)}</span>
      </S.DateWrapper>
      {todoList.length ? (
        <DragDropContext
          onDragEnd={(result) => {
            if (!result.destination) {
              return;
            }
            console.log(result);
            const updatedList = [...todoList];
            const [removed] = updatedList.splice(result.source.index, 1);
            updatedList.splice(result.destination.index, 0, removed);
            setTodoList(updatedList);
          }}
        >
          <Droppable droppableId="Todos">
            {(provided) => (
              <S.DraggableWrapper
                {...provided.droppableProps}
                ref={provided.innerRef}
              >
                {todoList.map((data: any, i) => (
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
      ) : (
        "Let's write todo!!"
      )}

      <Input setTodoList={setTodoList} todoList={todoList} />
    </S.Contain>
  );
};

export default TodoList;

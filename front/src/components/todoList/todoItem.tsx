import { useState } from "react";
import * as S from "./todoItem.style";
import axios from "axios";
const TodoItem = ({ todo, provided }) => {
  const { id, title, content } = todo;
  const [isChecked, setIsChecked] = useState(todo.isChecked);
  const togleChecked = async () => {
    await axios.post("http://localhost:8080/todo/toggle/" + id);
    setIsChecked(!isChecked);
    console.log(isChecked);
  };
  return (
    <S.Contain ref={provided.innerRef} {...provided.draggableProps}>
      <S.Wrap {...provided.dragHandleProps}>
        <S.Title>{title}</S.Title>
        <S.Content>{content}</S.Content>
      </S.Wrap>
      <S.CheckBox onClick={togleChecked} isChecked={isChecked}></S.CheckBox>
    </S.Contain>
  );
};

export default TodoItem;

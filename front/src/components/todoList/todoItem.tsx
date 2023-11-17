import { useState } from "react";
import * as S from "./todoItem.style";

const TodoItem = ({ todo }) => {
  const { id, title, content } = todo;
  const [isChecked, setIsChecked] = useState(todo.isChecked);
  const togleChecked = () => {
    setIsChecked(!isChecked);
    console.log(isChecked);
  };
  return (
    <S.Contain>
      <S.Title>{title}</S.Title>
      <S.Content>{content}</S.Content>
      <S.CheckBox onClick={togleChecked} isChecked={isChecked}></S.CheckBox>
    </S.Contain>
  );
};

export default TodoItem;

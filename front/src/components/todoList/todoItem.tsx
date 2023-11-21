import { useState } from "react";
import * as S from "./todoItem.style";
import axios from "axios";
const TodoItem = ({ todo }) => {
  const { id, title, content } = todo;
  const [isChecked, setIsChecked] = useState(todo.isChecked);
  const togleChecked = async () => {
    await axios.post("http://localhost:8080/todo/toggle/" + id);
    setIsChecked(!isChecked);
    console.log(isChecked);
  };
  return (
    <S.Contain>
      <S.Wrap>
        <S.Title>{title}</S.Title>
        <S.Content>{content}</S.Content>
      </S.Wrap>
      <S.CheckBox onClick={togleChecked} isChecked={isChecked}></S.CheckBox>
    </S.Contain>
  );
};

export default TodoItem;

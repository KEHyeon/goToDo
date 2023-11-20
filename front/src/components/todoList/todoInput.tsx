import * as S from "./todoInput.style.ts";

const TodoInput = () => {
  return (
    <S.Contain>
      INPUT
      <S.Form>
        <div>
          Todo
          <S.Input />
        </div>
        <div>
          Description
          <S.Input />
        </div>
        <S.Button>GO!</S.Button>
      </S.Form>
    </S.Contain>
  );
};

export default TodoInput;

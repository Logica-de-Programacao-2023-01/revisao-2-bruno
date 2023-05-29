function mergeStudentData(studentData1, studentData2) {
  let mergedData = {};

  // Itera sobre as chaves (nomes dos alunos) do mapa studentData1
  for (let name in studentData1) {
    if (name in studentData2) {
      // Se o nome do aluno também está presente no mapa studentData2, combina as informações
      mergedData[name] = {
        name: name,
        age: studentData1[name].age,
        subjects: { ...studentData1[name].subjects, ...studentData2[name].subjects }
      };
    } else {
      // Se o nome do aluno não está presente no mapa studentData2, copia as informações do mapa studentData1
      mergedData[name] = {
        name: name,
        age: studentData1[name].age,
        subjects: { ...studentData1[name].subjects }
      };
    }
  }

  // Itera sobre as chaves (nomes dos alunos) do mapa studentData2
  for (let name in studentData2) {
    if (!(name in studentData1)) {
      // Se o nome do aluno não está presente no mapa studentData1, copia as informações do mapa studentData2
      mergedData[name] = {
        name: name,
        age: studentData2[name].age,
        subjects: { ...studentData2[name].subjects }
      };
    }
  }

  return mergedData;
}

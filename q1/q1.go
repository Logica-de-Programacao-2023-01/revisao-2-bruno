function mergeStudentData(studentData1, studentData2) {
  let mergedData = {};

 
  for (let name in studentData1) {
    if (name in studentData2) {
     
      mergedData[name] = {
        name: name,
        age: studentData1[name].age,
        subjects: { ...studentData1[name].subjects, ...studentData2[name].subjects }
      };
    } else {
 
      mergedData[name] = {
        name: name,
        age: studentData1[name].age,
        subjects: { ...studentData1[name].subjects }
      };
    }
  }
  for (let name in studentData2) {
    if (!(name in studentData1)) 
      mergedData[name] = {
        name: name,
        age: studentData2[name].age,
        subjects: { ...studentData2[name].subjects }
      };
    }
  }

  return mergedData;
}

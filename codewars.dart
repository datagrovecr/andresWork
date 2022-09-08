// This function should test if the factor is a factor of base.
//Return true if it is a factor or false if it is not.
// link:
// https://www.codewars.com/kata/55cbc3586671f6aa070000fb/solutions/dart?filter=me&sort=best_practice&invalids=false
bool checkForFactor(int base, int factor) {
  num check = base % factor;
  if (check == 0) {
    return true;
  } else {
    return false;
  }
}

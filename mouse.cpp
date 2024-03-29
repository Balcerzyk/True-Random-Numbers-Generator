#include <iostream>
#include <winuser.h>
#include <windows.h>
#include <fstream>
#include <math.h>

int main()
{
    POINT p;
    double positionLabelX[129];
    double positionLabelY[129];
    /////////////////////////////////////////////// odczyt pozycji

    int x, y;
    std::cout<<"Poruszaj myszka"<<std::endl;
    for(int i=0; i<129; i++)
    {
        Sleep(20);
        GetCursorPos(&p);
        if(x==p.x && y==p.y)
            i--;
        else
        {
            x=p.x;
            y=p.y;
            positionLabelX[i]=x;
            positionLabelY[i]=y;
        }
    }

    /////////////////////////////////////////////// obliczenia
    double k[129];
    double r[129];
    for(int i=0; i<128; i++)
    {
        if(positionLabelX[i]==positionLabelX[i+1])
        {
            k[i]=1.57;
        }
        else
        {
            k[i]=atan((abs(positionLabelY[i+1]-positionLabelY[i])/abs(positionLabelX[i+1]-positionLabelX[i])));
        }
        r[i]=(k[i]/1.57);
        std::cout<<"x = "<<positionLabelX[i]<<"   y = "<<positionLabelY[i]<<"   kat = "<<k[i]<<"   r = "<<r[i]<<std::endl;
    }

    /////////////////////////////////////////////// zapis do pliku wynik�w
    std::ofstream myfile;
    myfile.open("position.txt");
    for(int i=0; i<128; i++)
    {
        myfile<<"x = "<<positionLabelX[i]<<"   y = "<<positionLabelY[i]<<"   kat = "<<k[i]<<"   r = "<<r[i]<<"\n";
    }
    myfile.close();
    std::cout<<std::endl<<"Wyniki zapisane w position.txt"<<std::endl;
    return 0;
}

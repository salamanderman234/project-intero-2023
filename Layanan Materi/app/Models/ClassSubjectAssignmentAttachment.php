<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\SoftDeletes;

class ClassSubjectAssignmentAttachment extends Model
{
    use HasFactory, SoftDeletes;
    protected $guarded = ["id"];

    public function class_subject_assignment()
    {
        return $this->belongsTo(ClassSubjectAssignment::class,'class_subject_assignment_id', 'id');
    }
}
